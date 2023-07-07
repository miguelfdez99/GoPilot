// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {backend} from '../models';

export function AddCronJob(arg1:string,arg2:string):Promise<void>;

export function Backup(arg1:backend.BackupOptions):Promise<void>;

export function CheckAdmin():Promise<boolean>;

export function CommandExists(arg1:string):Promise<boolean|string>;

export function CreateGroup(arg1:string,arg2:any):Promise<void>;

export function CreateUser(arg1:backend.User):Promise<void>;

export function DeleteGroup(arg1:string):Promise<void>;

export function DeleteUser(arg1:string,arg2:boolean,arg3:boolean):Promise<void>;

export function ExportLogs(arg1:string,arg2:string):Promise<void>;

export function FetchTrafficData():Promise<Array<backend.TrafficData>>;

export function GenerateKeys(arg1:string,arg2:string,arg3:string,arg4:boolean):Promise<string>;

export function GenerateSelfSignedCertificate(arg1:string):Promise<string>;

export function GetCPUUsage():Promise<Array<number>>;

export function GetDefaultPolicy(arg1:string):Promise<string>;

export function GetDiskUsage():Promise<backend.DiskUsage>;

export function GetDistribution():Promise<string>;

export function GetFirewallStatus():Promise<string>;

export function GetLSCPU():Promise<string>;

export function GetLogs(arg1:string,arg2:number):Promise<Array<string>>;

export function GetMemoryUsage():Promise<backend.MemoryUsage>;

export function GetNetworkInterfaces():Promise<Array<backend.NetworkInterface>>;

export function GetNetworkUsage():Promise<backend.NetworkUsage>;

export function GetProcessInfo():Promise<Array<backend.ProcessInfo>>;

export function GetSystemInfo():Promise<string>;

export function InstallPackage(arg1:string):Promise<void>;

export function ListCronJobs():Promise<Array<backend.CronJob>>;

export function ListPackages():Promise<Array<string>>;

export function ListUfw():Promise<Array<string>>;

export function ModifyGroup(arg1:string,arg2:number):Promise<void>;

export function ModifyUser(arg1:string,arg2:backend.User):Promise<void>;

export function RemoveAllCronJobs():Promise<void>;

export function RemoveCronJob(arg1:string):Promise<void>;

export function RemovePackage(arg1:string):Promise<void>;

export function SetDefaultPolicy(arg1:string,arg2:string):Promise<void>;

export function SetFirewallStatus(arg1:string):Promise<void>;

export function SetInterfaceStatus(arg1:string,arg2:string):Promise<void>;

export function TerminateProcess(arg1:number):Promise<void>;
