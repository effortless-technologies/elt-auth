FROM iron/go:dev

WORKDIR /app

ENV MONGO_ADDR='104.198.34.190:27017'

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/effortless-technologies/elt-auth
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get
RUN cd $SRC_DIR; go build -o api; cp api /app/

ENTRYPOINT ./api -mongoAddr=$MONGO_ADDR