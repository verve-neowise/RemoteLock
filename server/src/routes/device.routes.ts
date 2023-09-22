import { Router } from "express";
import { deleteDevice, deviceStatus, lockDevice, unlockDevice } from "../functions/device";

export default Router()
    .get('/status', deviceStatus)
    .get('/:id/lock', lockDevice)
    .get('/:id/unlock', unlockDevice)
    .get('/:id/delete', deleteDevice)
