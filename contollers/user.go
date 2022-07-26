package contollers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"githum.com/okumawat/go-getting-started/models"
)

type userContoller struct {
	userIdPattern *regexp.Regexp
}

func (uc userContoller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	fmt.Println(url)
	if uc.userIdPattern.MatchString(url) {
		matches := uc.userIdPattern.FindStringSubmatch(url)

		userId, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatal("Error converting userid")
		}
		user, _ := models.GetUserById(userId)
		json.NewEncoder(w).Encode(user)
	} else {
		w.Write([]byte(url))
	}

}

func newUserController() *userContoller {
	return &userContoller{userIdPattern: regexp.MustCompile(`^/users/(\d+)`)}
}
