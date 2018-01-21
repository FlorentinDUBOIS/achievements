import { baseUri } from '../app.config'

export class Achievements {
  static find () {
    return fetch(`${baseUri}/achievements/`)
      .then(res => res.json())
  }

  static get (id) {
    return fetch(`${baseUri}/achievements/${id}`)
      .then(res => res.json())
  }
}
