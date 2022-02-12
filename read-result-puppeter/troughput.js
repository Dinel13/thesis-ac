const fs = require("fs");
const puppeteer = require("puppeteer");

const run = async () => {
  let dir = process.argv[2];
  try {
    let all = [];
    const browser = await puppeteer.launch({
      headless: true,
      args: ["--no-sandbox"],
    });
    const page = await browser.newPage();
    for (i = 1; i <= 10; i++) {
      let fullPath;
      // cek if dir conatin dir that name with ended .jtl
      if (fs.existsSync(dir + i + ".jtl/index.html")) {
        fullPath = dir + i + ".jtl/index.html";
      } else if (fs.existsSync(dir + i + "/index.html")) {
        fullPath = dir + i + "/index.html";
      } else {
        console.log(`file ke ${i} not found`);
        continue;
      }
      await page.goto("file://" + fullPath, {
        waitUntil: "load",
        timeout: 0,
      });
      // wait for the page to load
      await page.waitForTimeout(250);
      const error = await page.evaluate(() => {
        let ET = document.getElementById("top5ErrorsBySamplerTable");
        let EB = ET.getElementsByTagName("tbody");
        let Etr = EB[0].getElementsByTagName("tr");
        let Etd = Etr[0].getElementsByTagName("td");
        return Etd[3].innerText;
      });
      if (error != 0) {
        console.log("ada error");
        continue;
      }
      // read value of table id = "table" first row  in tbody and column 4
      const value = await page.evaluate(() => {
        let table = document.getElementById("statisticsTable");
        let tbody = table.getElementsByTagName("tbody");
        let tr = tbody[0].getElementsByTagName("tr");
        let td = tr[0].getElementsByTagName("td");
        let td_value = td[11].innerText;
        return td_value;
      });
      console.log(value);
      all.push(parseFloat(value));
    }
    await browser.close();
    // log the average
    console.log(all.reduce((a, b) => a + b, 0) / all.length);
    console.log(dir + " is done");
  } catch (error) {
    console.error(error);
  }
};

run();
