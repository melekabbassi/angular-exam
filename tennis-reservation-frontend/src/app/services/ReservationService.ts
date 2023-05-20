import { Injectable } from '@angular/core';
import { Reservation } from '../interfaces/reservations';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ReservationService {
  private _reservations = new BehaviorSubject<Reservation[]>(this.getReservationsFromLocal());

  get reservations() {
    return this._reservations.asObservable();
  }

  saveReservationToLocal(reservation: Reservation) {
    const reservations = this.getReservationsFromLocal();
    reservations.push(reservation);
    localStorage.setItem('reservations', JSON.stringify(reservations));
    this._reservations.next(reservations); // Emit the new reservations list
  }

  getReservationsFromLocal(): Reservation[] {
    const reservations = localStorage.getItem('reservations');
    return reservations ? JSON.parse(reservations) : [];
  }

  saveAllReservations(reservations: Reservation[]) {
    localStorage.setItem('reservations', JSON.stringify(reservations));
  }
}
