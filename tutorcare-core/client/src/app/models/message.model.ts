import { User } from "./user.model";

export interface Message {
  sender_id?: string
  message_id?: number
  chatroom_id?: number
  message?: string
  is_deleted?: boolean
  timestamp?: string
  sender?: User
}
