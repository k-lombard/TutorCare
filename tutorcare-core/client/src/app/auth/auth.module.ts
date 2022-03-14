import {ModuleWithProviders, NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {LoginComponent} from './login/login.component';
import {MatCardModule} from "@angular/material/card";
import { MatInputModule } from "@angular/material/input";
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatCheckboxModule} from '@angular/material/checkbox';
import { NgxSliderModule } from '@angular-slider/ngx-slider';
import {MatTooltipModule} from '@angular/material/tooltip';
import {RouterModule} from "@angular/router";
import {ReactiveFormsModule, FormsModule} from "@angular/forms";
import {MatButtonModule} from "@angular/material/button";
import { StoreModule } from '@ngrx/store';
import {AuthService} from "./auth.service";
import * as fromAuth from './auth.reducer';
import {AuthGuard} from './auth.guard';
import { AuthEffects } from './auth.effects';
import { EffectsModule } from '@ngrx/effects';
import { NavbarComponent } from '../navbar/navbar.component';
import { SignupComponent } from '../signup/signup.component';
import { HomeComponent } from '../home/home.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { HttpModule } from '@angular/http';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatGridListModule } from '@angular/material/grid-list';
import {MatSelectModule} from '@angular/material/select';
import { GoogleMapsModule } from '@angular/google-maps'
import { FindCareComponent } from '../find-care/find-care.component';
import { BarRatingModule } from "ngx-bar-rating";
import { EditProfileComponent } from '../profile/edit-profile/edit-profile.component';
import { AuthInterceptor } from './auth.interceptor';
import { ToastrModule } from 'ngx-toastr';
import { CreateJobDialog, FindJobsComponent } from '../find-jobs/find-jobs.component';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {NgxMaterialTimepickerModule} from 'ngx-material-timepicker';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatNativeDateModule } from '@angular/material/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { NgxMatDatetimePickerModule, NgxMatNativeDateModule, NgxMatTimepickerModule } from '@angular-material-components/datetime-picker';
import { ApplyJobDialog } from '../find-jobs/apply-job/apply-job.component';
import { ApplicationsReceivedComponent } from '../find-jobs/applications-received/applications-received.component';
import { ActiveJobsComponent } from '../find-jobs/active-jobs/active-jobs.component';
import { EditJobDialog, MyJobPostingsComponent } from '../find-jobs/my-job-postings/my-job-postings.component';
import { ChatroomsComponent } from '../find-jobs/chatrooms/chatrooms.component';
import { VerifyComponent } from '../signup/verify/verify.component';
import { SimplebarAngularModule } from 'simplebar-angular';
import { NgSelectModule } from '@ng-select/ng-select';
import { ProfileComponent } from '../profile/profile.component';
import { AppliedToComponent } from '../find-jobs/applied-to/applied-to.component';
import { SidebarComponent } from '../find-jobs/sidebar/sidebar.component';
import { JobComponent } from '../job/job.component';
import { HistoryComponent } from '../find-jobs/history/history.component';

@NgModule({
    imports: [
        CommonModule,
        FormsModule,
        MatSelectModule,
        MatTooltipModule,
        ReactiveFormsModule,
        MatInputModule,
        MatCardModule,
        NgxSliderModule,
        MatSlideToggleModule,
        MatCheckboxModule,
        MatTooltipModule,
        HttpClientModule,
        MatButtonModule,
        ReactiveFormsModule,
        NgxMaterialTimepickerModule,
        HttpModule,
        NgSelectModule,
        SimplebarAngularModule,
        MatToolbarModule,
        MatSidenavModule,
        NgxMatDatetimePickerModule,
        NgxMatTimepickerModule,
        MatListModule,
        MatIconModule,
        BarRatingModule,
        MatGridListModule,
        RouterModule,
        GoogleMapsModule,
        MatDatepickerModule,
        MatNativeDateModule,
        NgxMatNativeDateModule,
        MatDialogModule,
        ToastrModule,
        MatFormFieldModule,
        StoreModule.forFeature('auth', fromAuth.authReducer),
        EffectsModule.forFeature([AuthEffects]),

    ],
    declarations: [
        LoginComponent,
        NavbarComponent,
        SidebarComponent,
        SignupComponent,
        HomeComponent,
        FindCareComponent,
        EditProfileComponent,
        FindJobsComponent,
        CreateJobDialog,
        ApplyJobDialog,
        ApplicationsReceivedComponent,
        ActiveJobsComponent,
        MyJobPostingsComponent,
        ChatroomsComponent,
        VerifyComponent,
        ProfileComponent,
        AppliedToComponent,
        EditJobDialog,
        JobComponent,
        HistoryComponent
    ],
    exports: [LoginComponent, NavbarComponent, ],
    entryComponents: [CreateJobDialog]
})
export class AuthModule {
    static forRoot(): ModuleWithProviders<AuthModule> {
        return {
            ngModule: AuthModule,
            providers: [
              AuthService,
              AuthGuard,
              { provide: MAT_DIALOG_DATA, useValue: {} },
              {
                provide: MatDialogRef,
                useValue: {}
              },
              {
                provide : HTTP_INTERCEPTORS,
                useClass: AuthInterceptor,
                multi   : true,
              }
            ]
        }
    }
}
