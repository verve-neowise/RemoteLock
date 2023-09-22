const devicesService = require('../services/devices.service')
const moment = require('moment')

module.exports = (req, res) => {
    console.log(devicesService.getAllDevices())
    res.render('devices', {
        moment: moment,
        devices: devicesService.getAllDevices(),
        error: req.flash()['error']
    })
}