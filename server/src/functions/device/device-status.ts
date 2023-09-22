import { Request, Response } from "express";
import DeviceService from "../../data/devices.service";

export default async (req: Request, res: Response) => {
    try {
        const { id, model } = req.query

        const exists = await DeviceService.findByIdAndModel(String(id), String(model))
    
        if (!exists) {
            return res.send('Not found')
        }
        
        await DeviceService.update(exists.id, {
            lastSync: new Date()
        })
    
        res.send(exists.status)
    }
    catch(e: any) {
        console.error(e)
        res.status(500).send('Server error. See logs for more')
    }
}