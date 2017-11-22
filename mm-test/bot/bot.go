package bot
import (
	mm "github.com/mattermost/platform/model"
)

func Connect(url, user, pass string) (*mm.Client4,  error) {
	client := mm.NewAPIv4Client(url)
	_, errPing := client.GetPing()
	if (errPing.Error != nil) {return client, errPing.Error}
	_, errLogin := client.Login(user, pass)
	return client, errLogin.Error
}
