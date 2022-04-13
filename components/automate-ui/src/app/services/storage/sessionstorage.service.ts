import { Injectable } from '@angular/core';
import { Storage } from './storage';
import { ChefSessionService } from '../chef-session/chef-session.service';

@Injectable()
export class SessionStorageService extends Storage {

  constructor(protected chefSession: ChefSessionService) { super(chefSession); }

  protected store(key: string, value: string): string {
    sessionStorage.setItem(this.prefixKey(key), value);

    return value;
  }

  protected retrieve(key: string): string | null {
    return sessionStorage.getItem(this.prefixKey(key));
  }

}
