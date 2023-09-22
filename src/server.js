const express = require('express')
const cors = require('cors')
const session = require('express-session');
const cookieParser = require('cookie-parser');
const flash = require('connect-flash');

const devicesPage = require('./handlers/devicesPage') 
const addDevice = require('./handlers/addDevice') 
const lockDevice = require('./handlers/lockDevice') 
const unlockDevice = require('./handlers/unlockDevice') 
const getStatus = require('./handlers/getDeviceStatus') 

const app = express()

app.use(cors())
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.use(cookieParser('secret'));
app.use(session({
    cookie: { maxAge: 60000 }, 
    resave: true, 
    saveUninitialized: false, 
    secret: 'secret'
}));

app.use(flash());

app.set('view engine', 'ejs');

app.get('/', devicesPage)
app.post('/', addDevice)
app.get('/lock/:id', lockDevice)
app.get('/unlock/:id', unlockDevice)
app.get('/status', getStatus)

app.listen(process.env.PORT || 3000, () => console.log('Server is running'))