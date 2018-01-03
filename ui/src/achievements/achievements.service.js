import Dexie from 'dexie'

const db = new Dexie('achievements')

export class Achievements {
  constructor () {
    db.version(1).stores({
      achievements: 'name,description'
    })
  }

  connect () {
    return db.open()
  }

  find () {
    return db.achievements.toArray()
  }

  insert ({ name, description }) {
    return db.achievements.add({ name, description })
  }
}
