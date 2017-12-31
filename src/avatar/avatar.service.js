import Identicon from 'identicon.js'
import md5 from 'md5'

export class Avatar {
  constructor (hash) {
    this.hash = md5(hash)
  }

  generate () {
    const icon = new Identicon(this.hash, {
      background: [255, 255, 255, 0], // rgba white
      margin: 0.1,
      size: 256,
      format: 'svg'
    })

    return `data:image/svg+xml;base64,${icon.toString()}`
  }
}
