package persistence

import (
	"aos/secret"
	"fmt"

	"github.com/go-redis/redis"
)

// DAO的实现
type secretDAO struct {
	client *redis.Client
}

func NewSecretDAO(client *redis.Client) secret.SecretDAO {
	dao := &secretDAO{
		client: client,
	}
	return dao
}

func (d *secretDAO) Add(secret secret.Secret) error {
	key := fmt.Sprintf("secret_%s", secret.AccessKey)
	err := d.client.HSet(key, "key", secret.AccessKey).Err()
	if nil != err {
		return err
	}
	err = d.client.HSet(key, "secret", secret.AccessSecret).Err()
	if nil != err {
		return err
	}
	return nil
}

func (d *secretDAO) FindOne(secretAccessKey string) (*secret.Secret, error) {
	fmt.Println("SecretDAORedis.FindOne")
	key := fmt.Sprintf("secret_%s", secretAccessKey)
	fmt.Println(key)
	hget := d.client.HGet(key, "secret")
	err := hget.Err()
	if nil != err {
		return nil, err
	}
	found := &secret.Secret{
		key,
		hget.Val(),
	}
	fmt.Println(found)
	return found, nil

}
