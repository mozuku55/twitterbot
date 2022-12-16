package keys
import (
	"github.com/ChimeraCoder/anaconda"
	."./entity"
)

type keys = struct{
	apikey apikey
	accessToken accessToken
	bearerToken bearerToken
	apiSecretkey apiSecretkey
	accessTokenSecret accessTokenSecret
}

twitterkey := &keys{
	apiKey: "PXVdSbxjNOzvmD8UHW1iA61L6",
	accessToken: "1603781766887510016-ZrPU9jqE6eZN7YyKr13sKAp3IG7upc",
	bearerToken: "AAAAAAAAAAAAAAAAAAAAAAv5kQEAAAAATR7YmaGTI2rFaWjuZsV7yWL9fGg%3DKbEMw6xWm1gJNeWGl7mGLJga5xNEVuBi23RkrJIzWlACf5wkUn",
	apiSecretkey: "bHTUnMaf4v4Kxh7eVsqzRLHR1lELEEwUfKAuaQvJ8Vev1jwO3P",
	accessTokenSecret: "mKkpMn0je8r13RSadbyDgxnpjDqNLDa3eSzUn6PtcYKdf"
}

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(keys.apiKey)
	anaconda.SetConsumerSecret(keys.apiKey)
	api := anaconda.NewTwitterApi(keys.accessToken, keys.accessToken)
	return api
}
