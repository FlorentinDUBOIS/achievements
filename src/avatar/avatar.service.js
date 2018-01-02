import Identicon from 'identicon.js'
import md5 from 'md5'

export class Avatar {
  constructor (hash) {
    this.hash = md5(hash)
  }

  generate () {
    const icon = new Identicon(this.hash, {
      margin: 0.2,
      size: 256,
      format: 'svg'
    })

    return `data:image/svg+xml;base64,${icon.toString()}`
  }
}
