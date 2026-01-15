package pokeapi
import(
	"encoding/json"
	"io"
	"net/http"
	"errors"
)

//GetPokemon

func (c *Client) GetPokemon(name string) (Pokemon, error){
	url:= baseURL + "/pokemon/" + name 
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err!= nil{
			return pokemon, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err!= nil{
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Pokemon{}, err
	}

	if resp.StatusCode == 404{
		return Pokemon{}, errors.New("Pokemon does not exist")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err!=nil{
		return pokemon, nil
	}

	c.cache.Add(url, data)
	return pokemon, nil
}
