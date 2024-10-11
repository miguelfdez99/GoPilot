export namespace backend {
	
	export class BackupOptions {
	    SourceDir: string;
	    DestDir: string;
	
	    static createFrom(source: any = {}) {
	        return new BackupOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SourceDir = source["SourceDir"];
	        this.DestDir = source["DestDir"];
	    }
	}
	export class CronJob {
	    Schedule: string;
	    Command: string;
	
	    static createFrom(source: any = {}) {
	        return new CronJob(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Schedule = source["Schedule"];
	        this.Command = source["Command"];
	    }
	}
	export class DiskUsage {
	    Total: number;
	    Used: number;
	
	    static createFrom(source: any = {}) {
	        return new DiskUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Total = source["Total"];
	        this.Used = source["Used"];
	    }
	}
	export class MemoryUsage {
	    ram: number;
	    swap: number;
	
	    static createFrom(source: any = {}) {
	        return new MemoryUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ram = source["ram"];
	        this.swap = source["swap"];
	    }
	}
	export class NetworkUsage {
	    bytes_sent: number;
	    bytes_received: number;
	
	    static createFrom(source: any = {}) {
	        return new NetworkUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bytes_sent = source["bytes_sent"];
	        this.bytes_received = source["bytes_received"];
	    }
	}
	export class ProcessInfo {
	    User: string;
	    Pid: number;
	    CpuPercent: number;
	    MemPercent: number;
	    VMS: number;
	    RSS: number;
	    TTY: string;
	    Status: string;
	    StartTime: number;
	    Cmdline: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.User = source["User"];
	        this.Pid = source["Pid"];
	        this.CpuPercent = source["CpuPercent"];
	        this.MemPercent = source["MemPercent"];
	        this.VMS = source["VMS"];
	        this.RSS = source["RSS"];
	        this.TTY = source["TTY"];
	        this.Status = source["Status"];
	        this.StartTime = source["StartTime"];
	        this.Cmdline = source["Cmdline"];
	    }
	}
	export class RunningService {
	    Name: string;
	    ActiveState: string;
	    Description: string;
	
	    static createFrom(source: any = {}) {
	        return new RunningService(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.ActiveState = source["ActiveState"];
	        this.Description = source["Description"];
	    }
	}
	export class Service {
	    Name: string;
	    StartupStatus: string;
	
	    static createFrom(source: any = {}) {
	        return new Service(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.StartupStatus = source["StartupStatus"];
	    }
	}
	export class SystemStatThresholds {
	    CPU: number;
	    Memory: number;
	    Disk: number;
	
	    static createFrom(source: any = {}) {
	        return new SystemStatThresholds(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CPU = source["CPU"];
	        this.Memory = source["Memory"];
	        this.Disk = source["Disk"];
	    }
	}
	export class User {
	    Username: string;
	    Password: string;
	    UID: number;
	    GID: number;
	    HomeDirectory: string;
	    Shell: string;
	    Groups: string[];
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Username = source["Username"];
	        this.Password = source["Password"];
	        this.UID = source["UID"];
	        this.GID = source["GID"];
	        this.HomeDirectory = source["HomeDirectory"];
	        this.Shell = source["Shell"];
	        this.Groups = source["Groups"];
	    }
	}

}

