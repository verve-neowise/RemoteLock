import { Device } from '@prisma/client'
import prisma from './prisma'

interface CreateInput {
    deviceId: string,
    model: string,
    status: string,
}

interface UpdateInput {
    deviceId?: string,
    model?: string,
    status?: string,
    lastSync?: Date
}

export default class DeviceService {

    static fetchAll(): Promise<Device[]> {
        return prisma.device.findMany({})
    }

    static create(input: CreateInput): Promise<Device> {
        return prisma.device.create({
            data: {
                deviceId: input.deviceId,
                model: input.model,
                lastSync: new Date(),
                status: input.status
            }
        })
    }

    static update(id: string, input: UpdateInput): Promise<Device> {
        return prisma.device.update({
            where: {
                id
            },
            data: {
                deviceId: input.deviceId,
                model: input.model,
                status: input.status,
                lastSync: input.lastSync
            }
        })
    }

    static delete(id: string): Promise<Device>  {
        return prisma.device.delete({
            where: {
                id
            }
        })
    }

    static findById(id: string): Promise<Device | null>  {
        return prisma.device.findUnique({
            where: {
                id
            }
        })
    }

    static findByIdAndModel(deviceId: string, model: string): Promise<Device | null> {
        return prisma.device.findFirst({
            where: {
                AND: [
                    { deviceId }, { model }
                ]
            }
        })
    }
}    