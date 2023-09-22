import { Router } from "express";
import { getDevices, postDevice } from "../functions/main";

export default Router()
    .get('/', getDevices)
    .post('/', postDevice)