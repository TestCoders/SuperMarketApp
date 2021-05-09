﻿using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Service.Interfaces;
using Service.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace WebApplication.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ProvisionController : ControllerBase
    {
        private readonly ILogger<ProvisionController> _logger;
        private readonly ILijpeVoorraadServerService _voorraadService;

        public ProvisionController(ILogger<ProvisionController> logger, ILijpeVoorraadServerService voorraadService)
        {
            _voorraadService = voorraadService;
            _logger = logger;
        }

        [HttpPost]
        public async Task<ActionResult> PostSupply(ProvisioningRequest provision)
        {
            _voorraadService.PostSupply(provision);

            return new OkObjectResult(provision);
        }
    }
}
