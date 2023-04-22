// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function AddCronJob(arg1:string,arg2:string):Promise<void>;

export function AddFirewallRule(arg1:string):Promise<void>;

export function CheckAdmin():Promise<boolean>;

export function CreateGroup(arg1:string,arg2:any):Promise<void>;

export function CreateUser(arg1:main.User):Promise<void>;

export function DeleteGroup(arg1:string):Promise<void>;

export function DeleteUser(arg1:string,arg2:boolean,arg3:boolean):Promise<void>;

export function GetDistribution():Promise<string>;

export function GetLSCPU():Promise<string>;

export function GetSystemInfo():Promise<string>;

export function Greet(arg1:string):Promise<string>;

export function InstallPackage(arg1:string):Promise<void>;

export function ListCronJobs():Promise<Array<main.CronJob>>;

export function ListFirewallRules():Promise<Array<string>>;

export function ListPackages():Promise<Array<string>>;

export function ModifyGroup(arg1:string,arg2:number):Promise<void>;

export function ModifyUser(arg1:string,arg2:any):Promise<void>;

export function RemoveAllCronJobs():Promise<void>;

export function RemoveFirewallRule(arg1:string):Promise<void>;

export function RemovePackage(arg1:string):Promise<void>;
