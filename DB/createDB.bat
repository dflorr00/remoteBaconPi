@echo off
start cmd /k "cd %~dp0 && mysql -u root -padmin < tfg.sql"