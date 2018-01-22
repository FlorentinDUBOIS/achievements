import { baseUri } from '../app.config'

export class RankingService {
  static async get () {
    const res = await fetch(`${baseUri}/ranking/`)
    const results = await res.json()

    for (const idx in results) {
      const res = await fetch(`${baseUri}/users/${results[idx].user_id}`)
      const user = await res.json()

      results[idx] = {
        user,
        ...results[idx]
      }
    }

    return results.sort((a, b) => {
      if (a.count > b.count) return 1
      if (a.count < b.count) return -1
      return 0
    })
  }
}
