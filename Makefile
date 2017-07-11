DEPS := github.com/julienschmidt/httprouter
DEPS += github.com/jinzhu/gorm
DEPS +=	github.com/garyburd/redigo/redis
DEPS +=	github.com/lib/pq
DEPS +=	github.com/sirupsen/logrus
DEPS +=	gopkg.in/go-playground/validator.v9
DEPS +=	github.com/asaskevich/govalidator
DEPS +=	github.com/dgrijalva/jwt-go

BUILDNAME := services 

build: deps
	go build

deps:
	go get -v $(DEPS)

updatedeps:
	go get -v -u $(DEPS)

rundev: build
	./$(BUILDNAME) -debug=true
	
run:
	go build
	./$(BUILDNAME) -debug=false > log_$(BUILDNAME) 

clean:
	rm -f $(BUILDNAME)


