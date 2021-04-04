package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connection string for the db 

const connectionString = 'Connection String'

//database name 
const dbName = 'testing'

//name for the collection protocol 

const collName = 'todolist'

var collection *mongo.collection


