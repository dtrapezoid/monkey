/* Before you execute the program, Launch `cqlsh` and execute:
CREATE TABLE reviews_by_day
(
day int,
productid text,
 userid text, 
reviewid uuid,
profilename text, 
helpfulness text,
 score text, 
summary text, 
review  text, 
time timestamp,
PRIMARY KEY (reviewid, userid, productid, time)
)

*/
package main

import (
    "time"
    "strconv"
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "github.com/gocql/gocql"
)

func main() {
    // connect to the cluster
    cluster := gocql.NewCluster("localhost")
    cluster.Keyspace = "foods"
    cluster.Consistency = gocql.Quorum
    session, _ := cluster.CreateSession()
    defer session.Close()

    //read the file
    file, err := os.Open("finefoods.txt")
    if err != nil {
        panic(err)
    }
    
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var insertArray [8]string
    i := 0
    for scanner.Scan() {
    	thisLine := scanner.Text()
//	fmt.Println(thisLine + " with length " + strconv.Itoa(len(thisLine)))

	if (len(thisLine) > 0){
	    stringArray := strings.Split(thisLine, ":")
	
            fmt.Println(strings.TrimSpace(stringArray[0]))
	    fmt.Println(strings.TrimSpace(stringArray[1]))

	    insertArray[i] = strings.TrimSpace(stringArray[1]);
	    i = i+1

	} else {

	    seconds, err:= strconv.Atoi(insertArray[5])
	    if err != nil{
	   	fmt.Println(err);
	     }
	    days := seconds / 86400
    	    myTime := time.Unix(int64(seconds),0)
	    uuid, err := gocql.RandomUUID()

	    fmt.Println("LOOK HERE" , days,insertArray[0],insertArray[1],insertArray[2],uuid, insertArray[3],insertArray[4],insertArray[5],insertArray[6] ,myTime );

//insert a record
	    if err := session.Query(`INSERT INTO reviews_by_day (day, productid, userid, profilename, reviewid, helpfulness, score, summary, review, time) VALUES (?,?,?,?,?,?,?,?,?,?)`,
           days,insertArray[0],insertArray[1],insertArray[2],uuid, insertArray[3],insertArray[4],insertArray[6],insertArray[7] ,myTime).Exec(); err != nil {
        	log.Fatal(err)
	    }


	    i= 0
        }

    }




    var id gocql.UUID
    var text string

    /* Search for a specific set of records whose 'timeline' column matches
     * the value 'me'. The secondary index that we created earlier will be
     * used for optimizing the search */
    if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
        "me").Consistency(gocql.One).Scan(&id, &text); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Tweet:", id, text)

    // list all tweets
    iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
    for iter.Scan(&id, &text) {
        fmt.Println("Tweet:", id, text)
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
    }


}
