FROM golang

ARG app_env
ENV APP_ENV ${app_env}
 
WORKDIR /go/src/myproject
COPY . myproject/ 


RUN go get ./... 
RUN go build -o main .

CMD if [${app_env} = production]; \
        then \
        myproject; \ 
        else \ 
        go get github.com/pilu/fresh && \
        fresh; \
        fi

EXPOSE 9090
        