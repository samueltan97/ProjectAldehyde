using System.ComponentModel.DataAnnotations;

namespace NetCoreRESTTemplate.Models
{
    public class MemberRequest
    {
        [Required]
        public string Name { get; set; }
        [Required]
        public string Job { get; set; }
        [Required]
        public int Age { get; set; }
    }
}