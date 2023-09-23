import { Request, Response } from "express";
import DeviceService from "../../data/devices.service";

export default async (req: Request, res: Response) => {
    try {

        const { name, deviceId, model, status } = req.body

        const exists = await DeviceService.findById(deviceId)

        if (exists) {
            req.flash('error', `Device with id ${deviceId} already exists`)
            return res.redirect('/')
        }

        await DeviceService.create({
            name: name,
            deviceId: deviceId,
            model: model,
            status: status
        })

        res.redirect('/')
    }
    catch (e: any) {
        console.error(e)
        res.status(500).send('Server error. See logs for more')
    }
}