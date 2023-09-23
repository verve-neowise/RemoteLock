import { Request, Response } from "express";
import moment from "moment";
import DeviceService from "../../data/devices.service";

export default async (req: Request, res: Response) => {
    try {

        const devices = await DeviceService.fetchAll()

        res.render('devices', {
            moment: moment,
            devices: devices,
            error: req.flash()['error']
        })
    }
    catch(e: any) {
        console.error(e)
        res.status(500).send('Server error. See logs for more')
    }
}