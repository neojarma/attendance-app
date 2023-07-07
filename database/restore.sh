#!/bin/bash
/opt/mssql-tools/bin/sqlcmd -U sa -P Superadmin123@_ -Q "RESTORE DATABASE presensi FROM DISK='/var/opt/mssql/backup/backup.bak' WITH MOVE 'presensi' TO '/var/opt/mssql/data/presensi.mdf', MOVE 'presensi_log' TO '/var/opt/mssql/data/presensi_log.ldf', REPLACE;"
