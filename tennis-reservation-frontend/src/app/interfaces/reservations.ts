export interface Reservation {
    id: number;
    type: string;
    date: string; // or Date if you prefer
    hour: number;
    duration: number;
}
