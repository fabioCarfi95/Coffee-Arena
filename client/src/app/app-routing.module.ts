import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { HOME, LOGIN, QUIZ_ROOM, RANKING } from './constants/constants';

import { Page404Component } from './components/page404/page404.component';
import { LoginComponent } from './components/login/login.component';
import { HomeComponent } from './components/home/home.component';
import { QuizRoomComponent } from './components/quiz-room/quiz-room.component';
import { RankingComponent } from './components/ranking/ranking.component';

const routes: Routes = [
  {path: '', redirectTo: LOGIN, pathMatch: 'full'},
  {path: LOGIN, component: LoginComponent},
  {path: HOME, component: HomeComponent},
  {path: QUIZ_ROOM, component: QuizRoomComponent},
  {path: RANKING, component: RankingComponent},
  {path:"**", component:Page404Component}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

// Insert here all the components related to routing actions
export const routingComponents = [
  LoginComponent,
  HomeComponent,
  QuizRoomComponent,
  RankingComponent,
  Page404Component
]
