%undefine _missing_build_ids_terminate_build
%define _unpackaged_files_terminate_build 0
%global stg_dir     /etc/s3-code-deployer
%global bin_dir     %{stg_dir}/bin
%global init_dir    /etc/init.d
%global usr_dir     /usr/share
%define _topdir     /root/rpmbuild
%define name        s3-code-deployer
%define release     el%{rhel}
%define version     VERSION_NUMBER
%define buildroot   %{_topdir}/%{name}-%{version}


Name:           %{name}
Summary:        S3 CodeDeployer
License:        Private
Version:        %{version}
Release:        %{release}
Group:          Applications/Security
BuildArch:      x86_64
BuildRoot:      %{buildroot}
URL:            https://miradnan.github.io/s3-code-deployer/
Source0:        %{_topdir}/SOURCES/%{name}-%{version}.tar.gz

%description
AWS S3 CodeDeployer. Downloads and installed tar ball, zip files into your application directories.

%prep
%setup -c

%build
# nothing to build, its already a built file

%install
rm -rf $RPM_BUILD_ROOT

# stg dir
STG_PATH=$RPM_BUILD_ROOT/%{stg_dir}
if [ ! -d "$STG_PATH" ]; then
    mkdir -p $STG_PATH
fi
# bin dir
STG_PATH=$RPM_BUILD_ROOT/%{bin_dir}
if [ ! -d "$STG_PATH" ]; then
    mkdir -p $STG_PATH
fi

# init dir
STG_PATH=$RPM_BUILD_ROOT/%{init_dir}
if [ ! -d "$STG_PATH" ]; then
    mkdir -p $STG_PATH
fi

if [ ! -d "/var/log/s3-code-deployer" ]; then
    mkdir /var/log/s3-code-deployer
fi

if [ ! -f "/var/log/s3-code-deployer/agent.log" ]; then
    touch /var/log/s3-code-deployer/agent.log && touch /var/log/s3-code-deployer/agent.err
else
    echo "" > /var/log/s3-code-deployer/agent.log
    echo "" > /var/log/s3-code-deployer/agent.err
fi

mkdir -p $RPM_BUILD_ROOT%{usr_dir}
# Service
mkdir -p $RPM_BUILD_ROOT/lib/systemd/system/

ls

cd /root/project/shared/rpm/package/

# The actual moving
cp etc/s3-code-deployer/bin/agent $RPM_BUILD_ROOT/etc/s3-code-deployer/bin/agent
# Service 
cp lib/systemd/system/s3-code-deployer.service $RPM_BUILD_ROOT/lib/systemd/system/s3-code-deployer.service
# osquery
cp etc/init.d/s3-code-deployer $RPM_BUILD_ROOT/etc/init.d/s3-code-deployer
#cp -R usr/share $RPM_BUILD_ROOT%{usr_dir}


%files
%defattr(-,root,root,-) 
%{bin_dir}/agent
/lib/systemd/system/%{name}.service
%{init_dir}/%{name}
#%{usr_dir}


%post
chmod +x %{bin_dir}/agent
chmod +x %{init_dir}/%{name}
chmod +x /lib/systemd/system/%{name}.service

# add service to services list
/sbin/chkconfig --add %{name}
/sbin/chkconfig %{name} on


if [ -f "/bin/systemctl" ]; then
    # el7
    /bin/systemctl enable s3-code-deployer.service
    /bin/systemctl restart s3-code-deployer.service
else
    # el6
    /sbin/service %{name} restart
fi


%preun
if [ "$1" = "0" ]; then
    /sbin/service %{name} stop
    /sbin/chkconfig %{name} off
fi

%postun
#last removal
#rm -rf /etc/s3-code-deployer
#rm -rf /var/log/s3-code-deployer
/sbin/chkconfig --del %{name}


%clean
rm -rf $RPM_BUILD_ROOT

%changelog
    * Sun Oct 7 2018 Mir Adnan <miradnanali@gmail.com>
    - Initial version
