## Bitnine Coding Test Q3
### Driver (Go)
### Source code [[Link](driver)]

### Development Enviroment
Linux Ubuntu 20.04

go version go1.13.8 linux/amd64


### Instructions
    1. Install `pq`
    2. Run script in terminal 


### Progress
I used pq library to retrieve data from database and encoding to marshal it as json. I am able to retrieve data from database and correctly parse it in a Go struct. The marshaling of struct to json is resulting in empty output and I could not debug further. As in the screenshot below, the first three lines indicate the correctly parsed output, the last line is the json marshal result which I could not debug further.


### Output Sample
![JSON-output](./images/Screenshot%20from%202023-01-17%2000-24-26.png)