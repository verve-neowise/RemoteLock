import { Request, Response } from "express";
import DeviceService from "../../data/devices.service";

export default async (req: Request, res: Response) => {
    try {
        const deviceId = req.params.id
        const device = await DeviceService.findById(deviceId)
        
        if (!device) {
            req.flash('error', `Device not found ${deviceId}`)
            return res.redirect('/')
        }
    
        const updated = await DeviceService.update(deviceId, {
            status: 'unlocked'
        })

        res.redirect('/')
    }
    catch(e: any) {
        console.error(e)
        res.status(500).send('Server error. See logs for more')
    }
}