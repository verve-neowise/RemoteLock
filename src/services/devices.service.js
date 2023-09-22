const fs = require('fs')

let devices = []

const getAllDevices = () => {
    loadChanges()
    return devices
}

const createDevice = (device) => {
    devices.push(device)
    saveChanges()
}

const updateDevice = (id, device) => {
    const index = devices.findIndex(it => it.id == id)
    if (index != -1) {
        devices[index] = device
    }
    saveChanges()
} 

const removeDevice = (id) => {
    const index = devices.findIndex(it => it.id == id)
    if (index != -1) {
        devices.splice(index, 1)
    }
    saveChanges()
}

const findDeviceById = (id) => {
    return devices.find(it => it.id == id)
}

const findDeviceByIdAndModel = (id, model) => {
    loadChanges()
    return devices.find(it => it.id == id && it.model == model)
}

const saveChanges = () => {
    fs.writeFileSync('data.json', JSON.stringify(devices), {
        encoding: 'utf-8'
    })
}

const loadChanges = () => {

    if (!fs.existsSync('data.json')) {
        return
    }

    const content = fs.readFileSync('data.json', 'utf-8')
    devices = JSON.parse(content)
}

module.exports = {
    getAllDevices,
    createDevice,
    updateDevice,
    removeDevice,
    findDeviceById,
    findDeviceByIdAndModel
}