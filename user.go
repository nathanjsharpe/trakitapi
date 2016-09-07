package trakitapi

import "net/http"

const userBasePath = "user"

type UserService interface {
	List() ([]User, *http.Response, error)
}

type User struct {
	Id          int           `json:"uuid"`
	Username    string        `json:"username"`
	Tenant      string        `json:"tenant"`
	Permissions []interface{} `json:"permissions"`
	Data        UserData      `json:"data"`
}

type UserData struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Alias            string `json:"alias"`
	Description      string `json:"description"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Group            string `json:"group"`
	RoleId           int    `json:"role_id"`
	FailedLoginCount []int  `json:"failedLoginCount"`
	LastUpdate       string `json:"last_update"`
}

type UserServiceOp struct {
	client *Client
}

func (s *UserServiceOp) List() ([]User, *http.Response, error) {
	req, err := s.client.NewRequest("GET", userBasePath, nil)

	if err != nil {
		return nil, nil, err
	}

	users := make([]User, 0)
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, err
}

// func GetUser(id int) (User, error) {
// 	var user User

// 	res := Get("user/" + strconv.Itoa(id))

// 	err := json.NewDecoder(res.Body).Decode(&user)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	fmt.Println(user)

// 	return user, nil
// }
