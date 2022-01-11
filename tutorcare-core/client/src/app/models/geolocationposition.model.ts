import { User } from './user.model';

export interface GeolocationPositionWithUser {
    user_id: string,
    location_id: number,
    accuracy: number,
    latitude: number,
    longitude: number,
    user: User
}