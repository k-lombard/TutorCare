import { User } from "./user.model";

export interface Chatroom {
  user1_id?: string
  user2_id?: string
  chatroom_id?: number
  is_deleted?: boolean
  date_created?: string
  user1?: User
  user2?: User
}
