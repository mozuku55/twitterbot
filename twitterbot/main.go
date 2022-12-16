package main

import (
"fmt"
"time"
. "fmt"
. "./keys"
. "github.com/***/tweetbot/text" //text.goをgithubにあげて、go getをした
)

func main() {
    api :=  GetTwitterApi();
    fmt.println("done\n");
}