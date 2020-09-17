import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { TeslaCardsComponent } from './component/tesla-cards/tesla-cards.component';


const routes: Routes = [{path: '', pathMatch: 'full', component: TeslaCardsComponent}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
