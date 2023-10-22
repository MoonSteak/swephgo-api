
## Setup environment

### Build swisseph library (for run locally)
```
git clone https://github.com/aloistr/swisseph.git
cd swisseph
make libswe.so
cp libswe.so /usr/local/lib/
```

### Download Nasa JPL ephemeris files
- [DE430](https://ssd.jpl.nasa.gov/ftp/eph/planets/Linux/de430/)
- [DE431](https://ssd.jpl.nasa.gov/ftp/eph/planets/Linux/de431/)
- [DE440](https://ssd.jpl.nasa.gov/ftp/eph/planets/Linux/de440/)
- [DE441](https://ssd.jpl.nasa.gov/ftp/eph/planets/Linux/de441/)

Detail information about ephemeris could be found on [JPL Planetary and Lunar Ephemerides](https://ssd.jpl.nasa.gov/planets/eph_export.html) page

```sh
wget -O jpl/de440.eph https://ssd.jpl.nasa.gov/ftp/eph/planets/Linux/de440/linux_p1550p2650.440
```

## Run in Docker

```sh
docker build . -t swephgo-api 
docker run -it -e EPHE_PATH=/jpl -e JPL_FILE=de440.eph -v /Users/vzverev/nenadev/nena-dev/swephgo-api/jpl/de440.eph:/jpl/de440.eph -p 127.0.0.1:3000:3000 --rm swephgo-api /app

```