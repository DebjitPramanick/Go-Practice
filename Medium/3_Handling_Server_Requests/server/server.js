const express = require('express')


const app = express()
app.use(express.json())
const port = 8000;
// http://localhost:8000

app.listen(port, () => {
    console.log(`Server is running at port: ${port}`)
})

app.get("/", (req, res) => {
    console.log("Server is running.")
    res.status(200).json({message: "Server is running."})
})

app.get("/read", (req, res) => {
    console.info(`Client called /read at ${new Date().getTime()}`)
    res.status(200).json({
        name: "Debjit",
        age: 21
    })
})

app.post("/create", (req, res) => {
    console.info(`Client called /create at ${new Date().getTime()}`)
    const body = req.body;
    res.status(200).json(body)
})