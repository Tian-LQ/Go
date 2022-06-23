package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

type UserUpdate struct {
	Email    string
	Username string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashed string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	return err == nil
}

type UserRepo interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur      UserRepo
	pr      ProfileRepo
	jwtConf *conf.JWT
	log     *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwtConf *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtConf: jwtConf, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwtConf.GetToken(), username)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	user := User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	err := uc.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.generateToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	user, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(user.PasswordHash, password) {
		return nil, errors.Unauthorized("user", "login failed")
	}
	return &UserLogin{
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    uc.generateToken(user.Username),
	}, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, username, email, password string) (*UserUpdate, error) {
	config := sarama.NewConfig()
	//设置
	//ack应答机制
	config.Producer.RequiredAcks = sarama.WaitForAll

	//发送分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	//回复确认
	config.Producer.Return.Successes = true

	// 添加生产者拦截器
	config.Producer.Interceptors = []sarama.ProducerInterceptor{&MyProducerInterceptor{str: "[producer]-"}}

	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "weatherStation"
	msg.Value = sarama.StringEncoder("test:weatherStation device")

	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
	}
	defer client.Close()

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return nil, err
	}
	fmt.Printf("pid:%v offset:%v \n ", pid, offset)

	return nil, nil
}

var wg sync.WaitGroup

func (uc *UserUsecase) GetProfile(ctx context.Context, username, email, password string) (*UserUpdate, error) {
	config := sarama.NewConfig()
	config.Consumer.Interceptors = []sarama.ConsumerInterceptor{&MyConsumerInterceptor{str: "-[consumer]"}}
	//创建新的消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("fail to start consumer", err)
	}
	//根据topic获取所有的分区列表
	partitionList, err := consumer.Partitions("weatherStation")
	if err != nil {
		fmt.Println("fail to get list of partition,err:", err)
	}
	fmt.Println(partitionList)
	//遍历所有的分区
	for p := range partitionList {
		//针对每一个分区创建一个对应分区的消费者
		pc, err := consumer.ConsumePartition("weatherStation", int32(p), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		//异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d Offse:%d Key:%v Value:%s \n",
					msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
	return nil, nil
}
