import express from 'express';
import cors from 'cors';
import session from 'express-session';
import cookieParser from 'cookie-parser';
import flash from 'connect-flash';
import mainRoutes from './routes/main.routes';
import deviceRoutes from './routes/device.routes';

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

app.use('/', mainRoutes)
app.use('/device', deviceRoutes)

app.listen(process.env.PORT || 3000, () => console.log('Server is running'))