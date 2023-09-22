const devicesService = require('../services/devices.service')

module.exports = (req, res) => {

    const { deviceId, model, status } = req.body

    const exists = devicesService.findDeviceById(deviceId)

    if (exists) {
        req.flash('error', `Device with id ${deviceId} already exists`)
        return res.redirect('/')
    }

    devicesService.createDevice({
        id: deviceId,
        model: model,
        lastSync: new Date(),
        status: status
    })

    res.redirect('/')
}