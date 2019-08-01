using System;
using System.ComponentModel.DataAnnotations;
using System.Runtime.InteropServices.ComTypes;
using Newtonsoft.Json.Linq;

namespace NetCoreRESTTemplate.Models
{
    public class Member
    {
        public Member(Guid id, string name, string job, int age)
        {
            this.Age = age;
            this.Id = id;
            this.Job = job;
            this.Name = name;
        }
        [Required]
        public Guid Id { get; set; }
        [Required]
        public string Name { get; set; }
        [Required]
        public string Job { get; set; }
        [Required]
        public int Age { get; set; }
    }
}
