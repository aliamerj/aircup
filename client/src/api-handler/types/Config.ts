interface Account {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
}

interface Disk {
  diskPath: string;
}

export interface ConfigRes {
  type: string;
  account: Account;
  disks: Disk[];
}
