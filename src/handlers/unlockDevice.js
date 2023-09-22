const devicesService = require('../services/devices.service')

module.exports = (req, res) => {
    const deviceId = req.params.id
    const device = devicesService.findDeviceById(deviceId)
    if (device) {
        device.status = 'unlocked'
    }
    
    devicesService.updateDevice(deviceId, device)

    res.redirect('/')
}