package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shriramsharma/grpc-tutorial/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	router := mux.NewRouter()

	router.HandleFunc("/add/{a}/{b}", func(writer http.ResponseWriter, request *http.Request) {
		context := request.Context()
		params := mux.Vars(request)
		a, err := strconv.ParseUint(params["a"], 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "{'error': 'Invalid Param %v'}", params["a"])
			return
		}
		b, err := strconv.ParseUint(params["b"], 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "{'error': 'Invalid Param %v'}", params["b"])
			return
		}

		req := &proto.Request{
			A: int64(a),
			B: int64(b),
		}

		resp, err := client.Add(context, req)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "%w", err)
			return
		}

		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, fmt.Sprint(resp.Result))

	}).Methods("GET")

	router.HandleFunc("/multiply/{a}/{b}", func(writer http.ResponseWriter, request *http.Request) {

	}).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to run server: %w", err)
	}

}
