import { Navigator } from '../navigator/navigator.service'

export class Camera {
  static async stream () {
    const source = await Navigator.getUserMedia({
      video: true
    })

    return URL.createObjectURL(source)
  }
}
