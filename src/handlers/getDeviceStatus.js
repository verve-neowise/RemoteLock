const devicesService = require('../services/devices.service')

module.exports = (req, res) => {

    const { id, model } = req.query

    console.log(id, model);

    const exists = devicesService.findDeviceByIdAndModel(id, model)

    if (!exists) {
        return res.send('Not found')
    }

    exists.lastSync = new Date()

    devicesService.updateDevice(id, exists)

    res.send(exists.status)
}