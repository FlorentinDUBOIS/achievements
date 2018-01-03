export class Navigator {
  static getUserMedia (constraints) {
    return new Promise((resolve, reject) => {
      navigator.getUserMedia(constraints, resolve, reject)
    })
  }
}
