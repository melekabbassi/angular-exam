import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { ReservationComponent } from './reservation/reservation.component';
import { ValidationComponent } from './validation/validation.component';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  { path: '', component: LoginComponent},
  { path: 'home', component: HomeComponent },
  { path: 'reservation', component: ReservationComponent },
  { path: 'validation', component: ValidationComponent },
  { path: 'signup', component: RegisterComponent },
  { path: 'signin', redirectTo: '' },
  { path: '**', redirectTo: '' }
];

// const routes: Routes = [
//   { path: 'signin', component: LoginComponent},
//   { path: 'home', component: HomeComponent },
//   { path: 'reservation', component: ReservationComponent },
//   { path: 'validation', component: ValidationComponent },
//   { path: 'signup', component: RegisterComponent },
//   { path: '', redirectTo: 'signin', pathMatch: 'full' },
//   { path: '**', redirectTo: 'signin' }
// ]

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
 