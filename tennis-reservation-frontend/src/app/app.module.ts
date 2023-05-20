import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { ReservationComponent } from './reservation/reservation.component';
import { HeaderComponent } from './header/header.component';
import { ValidationComponent } from './validation/validation.component';
import { RegisterComponent } from './register/register.component';
import { FooterComponent } from './footer/footer.component';
import { LoginComponent } from './login/login.component';
import { FormsModule, ReactiveFormsModule, } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { DashboardComponent } from './dashboard/dashboard.component';
import { HeaderdashboardComponent } from './headerdashboard/headerdashboard.component';
import { GroundsComponent } from './grounds/grounds.component';
import { ServicesTennisComponent } from './services-tennis/services-tennis.component';
import { MaterialsComponent } from './materials/materials.component';
import { ReservationsListComponent } from './reservations-list/reservations-list.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    ReservationComponent,
    HeaderComponent,
    ValidationComponent,
    RegisterComponent,
    FooterComponent,
    LoginComponent,
    DashboardComponent,
    HeaderdashboardComponent,
    GroundsComponent,
    ServicesTennisComponent,
    MaterialsComponent,
    ReservationsListComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
